package tihe.chen.springbootmongodb.controllers;

import tihe.chen.springbootmongodb.models.Media;
import tihe.chen.springbootmongodb.services.MediaService;
import io.swagger.v3.oas.annotations.Operation;
import io.swagger.v3.oas.annotations.security.SecurityRequirement;
import io.swagger.v3.oas.annotations.tags.Tag;
import org.springdoc.api.annotations.ParameterObject;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.Optional;

@RequestMapping("/api/medium")
@RestController
@Tag(name = "Media", description = "The Media API. Contains operations like create media, get media etc.")
public class MediaController {
    private final MediaService mediaService;

    public MediaController(MediaService mediaService) {
        this.mediaService = mediaService;
    }

    @Operation(summary = "Media Creation", description = "Create the media")
    @io.swagger.v3.oas.annotations.parameters.RequestBody(content = @io.swagger.v3.oas.annotations.media.Content(mediaType = "application/json", examples = @io.swagger.v3.oas.annotations.media.ExampleObject(value = "{\n" + "  \"name\": \"The Times\",\n"
            + "  \"thumbnail\": \"https://www.thetimes.com/images/logo.png\"\n" + "}", summary = "Media Creation Example")))
    @SecurityRequirement(name = "Bearer Authentication")
    @PostMapping(produces = MediaType.APPLICATION_JSON_VALUE)
    public ResponseEntity<Media> saveOrUpdateMedia(@RequestBody Media media) {
        return ResponseEntity.ok().body(mediaService.saveOrUpdateMedia(media));
    }

    @Operation(summary = "Media Search", description = "Get the media by Id")
    @GetMapping(value = "/{id}", produces = MediaType.APPLICATION_JSON_VALUE)
    public ResponseEntity<Media> getMediaById(@PathVariable String id) {
        return ResponseEntity.ok().body(mediaService.getMediaById(id));
    }

    @Operation(summary = "Media Search", description = "Get all medium or get the media by name if name is not null")
    @GetMapping(value = "/", produces = MediaType.APPLICATION_JSON_VALUE)
    public ResponseEntity<Page<Media>> getMedium(@RequestParam Optional<String> name, @ParameterObject Pageable pageable) {
        return ResponseEntity.ok().body(mediaService.getMedium(name, pageable));
    }

    @Operation(summary = "Media Deletion", description = "Delete the media by id")
    @SecurityRequirement(name = "Bearer Authentication")
    @DeleteMapping(value = "/{id}", produces = MediaType.APPLICATION_JSON_VALUE)
    public ResponseEntity<Media> deleteMedium(@PathVariable String id) {
        return ResponseEntity.ok().body(mediaService.deleteMediaById(id));
    }
}
