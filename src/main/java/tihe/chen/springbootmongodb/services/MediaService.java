package tihe.chen.springbootmongodb.services;

import tihe.chen.springbootmongodb.models.Media;
import org.bson.types.ObjectId;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.mongodb.core.MongoTemplate;
import org.springframework.data.mongodb.core.query.Criteria;
import org.springframework.data.mongodb.core.query.Query;
import org.springframework.data.support.PageableExecutionUtils;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.Optional;

@Service
public class MediaService {
    private final MongoTemplate mongoTemplate;

    public MediaService(MongoTemplate mongoTemplate) {
        this.mongoTemplate = mongoTemplate;
    }

    public Media saveOrUpdateMedia(Media mediaDto) {
        Media media = mongoTemplate.findOne(new Query(Criteria.where("name").is(mediaDto.getName())), Media.class);
        if (media == null) {
            media = new Media(mediaDto.getName(), mediaDto.getThumbnail());
        } else {
            media.setThumbnail(mediaDto.getThumbnail());
        }
        return mongoTemplate.save(media);
    }

    public Media getMediaById(String id) {
        return mongoTemplate.findById(new ObjectId(id), Media.class);
    }

    public Page<Media> getMedium(Optional<String> name, Pageable pageable) {
        List<Media> mediaList;
        Query query = new Query().with(pageable);
        if (name.isPresent()) {
            mediaList = mongoTemplate.find(query.addCriteria(Criteria.where("name").is(name.get())), Media.class);
        } else {
            mediaList = mongoTemplate.find(query, Media.class);
        }
        return PageableExecutionUtils.getPage(
                mediaList,
                pageable,
                () -> mongoTemplate.count(Query.of(query).limit(-1).skip(-1), Media.class));
    }

    public Media deleteMediaById(String id) {
        return mongoTemplate.findAndRemove(new Query(Criteria.where("id").is(id)), Media.class);
    }
}
