package com.digitalibrary.catalogue.mapper;

import com.digitalibrary.catalogue.dto.language.CreateLanguageRequest;
import com.digitalibrary.catalogue.dto.language.LanguageResponse;
import com.digitalibrary.catalogue.dto.language.UpdateLanguageRequest;
import com.digitalibrary.catalogue.entity.Language;
import org.springframework.stereotype.Component;

@Component
public class LanguageMapper extends BaseMapper<Language, CreateLanguageRequest, UpdateLanguageRequest, LanguageResponse> {

    @Override
    public Language createRequestToEntity(CreateLanguageRequest dto) {
        Language entity = new Language();

        entity.setCode(dto.code());
        entity.setValue(dto.value());
        entity.setDescription(dto.description());

        return entity;
    }

    @Override
    public Language updateRequestToEntity(Language entity, UpdateLanguageRequest dto) {
        if(shouldUpdateHelper(dto.code())) {
            entity.setCode(dto.code());
        }
        if(shouldUpdateHelper(dto.value())) {
            entity.setValue(dto.value());
        }
        if(shouldUpdateHelper(dto.description())) {
            entity.setDescription(dto.description());
        }

        return entity;
    }

    @Override
    public LanguageResponse toResponse(Language entity) {
        return new LanguageResponse(
                entity.getId(),
                entity.getCode(),
                entity.getValue(),
                entity.getDescription(),
                entity.getCreatedAt(),
                entity.getUpdatedAt()
        );
    }

    private static boolean shouldUpdateHelper(String value) {
        return value != null && !value.isBlank();
    }
}
