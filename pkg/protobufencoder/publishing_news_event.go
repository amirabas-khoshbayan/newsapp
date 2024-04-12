package protobufencoder

import (
	"encoding/base64"
	"github.com/labstack/gommon/log"
	"google.golang.org/protobuf/proto"
	"newsapp/contract/goproto/publishing"
	"newsapp/entity"
	"newsapp/pkg/slice"
)

func EncodePublishingNewsPublishedEvent(pNews entity.PublishedNews) string {
	protoBufNews := publishing.PublishingNews{Category: string(pNews.Category), NewsIds: slice.MapFromUintToUint64(pNews.NewsIDs)}

	payload, err := proto.Marshal(&protoBufNews)
	if err != nil {
		log.Error("proto.Marshal(&protoBufNews) error :", err.Error())

		return ""
	}

	return base64.StdEncoding.EncodeToString(payload)
}

func DecodePublishingNewsPublishedEvent(data string) entity.PublishedNews {
	payload, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		log.Error("base64.StdEncoding.DecodeString(data) error :", err.Error())

		return entity.PublishedNews{}
	}

	protoBufNews := publishing.PublishingNews{}
	if err := proto.Unmarshal(payload, &protoBufNews); err != nil {
		log.Error("proto.Unmarshal(payload, &protoBufNews) error :", err.Error())

		return entity.PublishedNews{}
	}

	return entity.PublishedNews{
		Category: entity.Category(protoBufNews.Category),
		NewsIDs:  slice.MapFromUint64ToUint(protoBufNews.NewsIds),
	}
}
