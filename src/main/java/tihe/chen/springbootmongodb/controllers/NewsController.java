package tihe.chen.springbootmongodb.controllers;

import tihe.chen.springbootmongodb.dtos.NewsDto;
import tihe.chen.springbootmongodb.models.News;
import tihe.chen.springbootmongodb.services.NewsService;
import io.swagger.v3.oas.annotations.Operation;
import io.swagger.v3.oas.annotations.security.SecurityRequirement;
import io.swagger.v3.oas.annotations.tags.Tag;
import org.springdoc.api.annotations.ParameterObject;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@RequestMapping("/api/news")
@RestController
@CrossOrigin(origins = "${allowed.origin}", allowCredentials = "true")
@Tag(name = "News", description = "The News API. Contains operations like create news, get news etc.")
public class NewsController {
    private NewsService newsService;

    public NewsController(NewsService newsService) {
        this.newsService = newsService;
    }

    @Operation(summary = "News Creation", description = "Create the news")
    @io.swagger.v3.oas.annotations.parameters.RequestBody(content = @io.swagger.v3.oas.annotations.media.Content(mediaType = "application/json", examples = @io.swagger.v3.oas.annotations.media.ExampleObject(value = "{\n" + "  \"media\": \"Times\",\n"
            + "  \"description\": \"https://www.theprimedia.com/skin/images/logo.png\"\n" +
            "  \"mediaUrl\": \"https://www.theprimedia.com/skin/images/logo.png\"\n" + "}", summary = "News Creation Example")))
    @SecurityRequirement(name = "Bearer Authentication")
    @PostMapping(produces = MediaType.APPLICATION_JSON_VALUE)
    public ResponseEntity<News> saveOrUpdateNews(@RequestBody NewsDto newsDto) {
        return ResponseEntity.ok().body(newsService.saveOrUpdateNews(newsDto));
    }

    @Operation(summary = "News Search", description = "Get news by Id")
    @GetMapping(value = "/{id}", produces = MediaType.APPLICATION_JSON_VALUE)
    public ResponseEntity<News> getNewsById(@PathVariable String id) {
        return ResponseEntity.ok().body(newsService.getNewsById(id));
    }

    @Operation(summary = "News Search", description = "Get all news")
    @GetMapping(value = "/", produces = MediaType.APPLICATION_JSON_VALUE)
    public ResponseEntity<Page<News>> getNews(@ParameterObject Pageable pageable) {
        return ResponseEntity.ok().body(newsService.getNews(pageable));
    }

    @Operation(summary = "News Deletion", description = "Delete news by id")
    @SecurityRequirement(name = "Bearer Authentication")
    @DeleteMapping(value = "/{id}", produces = MediaType.APPLICATION_JSON_VALUE)
    public ResponseEntity<News> deleteNewsById(@PathVariable String id) {
        return ResponseEntity.ok().body(newsService.deleteNewsById(id));
    }
}
