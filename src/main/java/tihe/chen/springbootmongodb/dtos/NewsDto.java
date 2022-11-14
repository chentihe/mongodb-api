package tihe.chen.springbootmongodb.dtos;

import javax.validation.constraints.NotNull;
import java.time.LocalDateTime;

public class NewsDto {
    @NotNull
    private String media;
    @NotNull
    private String description;
    @NotNull
    private String newsUrl;
    @NotNull
    private LocalDateTime creationDate;

    public String getMedia() {
        return media;
    }

    public void setMedia(String media) {
        this.media = media;
    }

    public String getDescription() {
        return description;
    }

    public void setDescription(String description) {
        this.description = description;
    }

    public String getNewsUrl() {
        return newsUrl;
    }

    public void setNewsUrl(String newsUrl) {
        this.newsUrl = newsUrl;
    }

    public LocalDateTime getCreationDate() {
        return creationDate;
    }

    public void setCreationDate(LocalDateTime creationDate) {
        this.creationDate = creationDate;
    }
}
