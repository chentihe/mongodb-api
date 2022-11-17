package tihe.chen.springbootmongodb.services;

import tihe.chen.springbootmongodb.dtos.NewsDto;
import tihe.chen.springbootmongodb.models.Media;
import tihe.chen.springbootmongodb.models.News;
import org.bson.types.ObjectId;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.mongodb.core.MongoTemplate;
import org.springframework.data.mongodb.core.query.Criteria;
import org.springframework.data.mongodb.core.query.Query;
import org.springframework.data.support.PageableExecutionUtils;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class NewsService {
    private final MongoTemplate mongoTemplate;

    public NewsService(MongoTemplate mongoTemplate) {
        this.mongoTemplate = mongoTemplate;
    }

    public News saveOrUpdateNews(NewsDto newsDto) {
        Media media = mongoTemplate.findOne(new Query(Criteria.where("name").is(newsDto.getMedia())), Media.class);
        News news = mongoTemplate.findOne(new Query(Criteria.where("newsUrl").is(newsDto.getNewsUrl())), News.class);
        if (news == null) {
            news = new News(media, newsDto.getDescription(), newsDto.getNewsUrl());
        } else {
            news.setDescription(newsDto.getDescription());
            news.setMedia(media);
        }
        return mongoTemplate.save(news);
    }

    public News getNewsById(String id) {
        return mongoTemplate.findById(new ObjectId(id), News.class);
    }

    public Page<News> getNews(Pageable pageable) {
        Query query = new Query().with(pageable);
        List<News> newsList = mongoTemplate.find(query, News.class);
        return PageableExecutionUtils.getPage(
                newsList,
                pageable,
                () -> mongoTemplate.count(Query.of(query).limit(-1).skip(-1), News.class));
    }

    public News deleteNewsById(String id) {
        return mongoTemplate.findAndRemove(new Query(Criteria.where("id").is(id)), News.class);
    }
}
