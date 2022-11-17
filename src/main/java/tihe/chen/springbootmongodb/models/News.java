package tihe.chen.springbootmongodb.models;

import org.springframework.data.mongodb.core.mapping.Document;
import org.springframework.data.mongodb.core.mapping.DocumentReference;
import org.springframework.data.mongodb.core.mapping.MongoId;

import java.time.LocalDateTime;
import java.util.Objects;

@Document("News")
public class News {
    @MongoId
    private String id;
    @DocumentReference(lookup = "{ 'name' : ?#{#target} }")
    private Media media;
    private String description;
    private String mediaUrl;
    private LocalDateTime creationDate;

    public News(Media media, String description, String mediaUrl) {
        this.media = media;
        this.description = description;
        this.mediaUrl = mediaUrl;
        this.creationDate = LocalDateTime.now();
    }

    public String getId() {
        return id;
    }

    public void setId(String id) {
        this.id = id;
    }

    public Media getMedia() {
        return media;
    }

    public void setMedia(Media media) {
        this.media = media;
    }

    public String getDescription() {
        return description;
    }

    public void setDescription(String description) {
        this.description = description;
    }

    public String getMediaUrl() {
        return mediaUrl;
    }

    public void setMediaUrl(String mediaUrl) {
        this.mediaUrl = mediaUrl;
    }

    public LocalDateTime getCreationDate() {
        return creationDate;
    }

    public void setCreationDate(LocalDateTime creationDate) {
        this.creationDate = creationDate;
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        News news = (News) o;
        return Objects.equals(id, news.id) && media.equals(news.media) && description.equals(news.description) && mediaUrl.equals(news.mediaUrl) && creationDate.equals(news.creationDate);
    }

    @Override
    public int hashCode() {
        return Objects.hash(id, media, description, mediaUrl, creationDate);
    }

    @Override
    public String toString() {
        return "News{" +
                "id='" + id + '\'' +
                ", media=" + media +
                ", description='" + description + '\'' +
                ", mediaUrl='" + mediaUrl + '\'' +
                ", creationDate=" + creationDate +
                '}';
    }
}
