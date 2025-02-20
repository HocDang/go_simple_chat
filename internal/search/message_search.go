package search

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"chat-server/internal/domain/entities"
	"chat-server/internal/domain/searches"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/google/uuid"
)

type messageSearch struct {
	Client *elasticsearch.Client
	Index  string
}

func NewEsMessage(client *elasticsearch.Client, index string) searches.MessageSearch {
	return &messageSearch{Client: client, Index: index}
}

func (es messageSearch) IndexMessage(msg entities.Message) error {
	data, err := json.Marshal(msg)

	if err != nil {
		return err
	}

	req := esapi.IndexRequest{
		Index: es.Index,
		Body:  bytes.NewReader(data),
	}

	res, err := req.Do(context.Background(), es.Client)
	if err != nil {
		return fmt.Errorf("error performing request: %v", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		body, _ := io.ReadAll(res.Body)
		return fmt.Errorf("Elasticsearch response error: %s", body)
	}

	return nil
}

func (es messageSearch) SearchMessages(senderID uuid.UUID, receiverID uuid.UUID, keyword string) ([]entities.Message, error) {
	query := fmt.Sprintf(`
	{
		"query": {
			"bool": {
				"must": [
					{ "match": { "content": "%s" } },
					{ "bool": {
						"should": [
							{
								"bool": {
									"must": [
										{ "match": { "sender_id": "%s" } },
										{ "match": { "receiver_id": "%s" } }
									]
								}
							},
							{
								"bool": {
									"must": [
										{ "match": { "sender_id": "%s" } },
										{ "match": { "receiver_id": "%s" } }
									]
								}
							}
						]
					}}
				]
			}
		}
	}`, keyword, senderID.String(), receiverID.String(), receiverID.String(), senderID.String())

	req := esapi.SearchRequest{
		Index: []string{es.Index},
		Body:  strings.NewReader(query),
	}

	res, err := req.Do(context.Background(), es.Client)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("elasticsearch search error: %s", res.Status())
	}

	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	hits := result["hits"].(map[string]interface{})["hits"].([]interface{})
	var messages []entities.Message
	for _, hit := range hits {
		source := hit.(map[string]interface{})["_source"]
		data, _ := json.Marshal(source)
		var msg entities.Message
		_ = json.Unmarshal(data, &msg)
		messages = append(messages, msg)
	}

	return messages, nil
}
