package tihe.chen.springbootmongodb.configs;

import com.mongodb.ConnectionString;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.data.mongodb.core.MongoClientFactoryBean;
import org.springframework.data.mongodb.core.MongoTemplate;

@Configuration
public class MongoDBConfiguration {
    @Bean
    public MongoClientFactoryBean mongoClientFactoryBean(@Value("${mongo.uri}") String uri) throws Exception {
        MongoClientFactoryBean mongodb = new MongoClientFactoryBean();
        ConnectionString conn = new ConnectionString(uri);
        mongodb.setConnectionString(conn);
        return mongodb;
    }

    @Bean
    public MongoTemplate mongoTemplate(MongoClientFactoryBean mongoClientFactoryBean, @Value("${mongo.database}") String database) throws Exception {
        return new MongoTemplate(mongoClientFactoryBean.getObject(), database);
    }
}
