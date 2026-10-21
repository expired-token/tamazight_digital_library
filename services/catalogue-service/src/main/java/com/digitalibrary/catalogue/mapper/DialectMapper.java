package com.digitalibrary.catalogue.mapper;

import com.digitalibrary.catalogue.dto.dialect.CreateDialectRequest;
import com.digitalibrary.catalogue.dto.dialect.DialectResponse;
import com.digitalibrary.catalogue.dto.dialect.UpdateDialectRequest;
import com.digitalibrary.catalogue.entity.Dialect;
import org.springframework.stereotype.Component;

@Component
public class DialectMapper extends BaseMapper<Dialect, CreateDialectRequest, UpdateDialectRequest, DialectResponse> {
    @Override
    public Dialect createRequestToEntity(CreateDialectRequest dto) {
        Dialect entity = new Dialect();

        entity.setCode(dto.code());
        entity.setValue(dto.value());
        entity.setDescription(dto.description());

        return entity;
    }

    @Override
    public Dialect updateRequestToEntity(Dialect entity, UpdateDialectRequest dto) {
        if(dto.code() != null && !dto.code().isBlank()) {
            entity.setCode(dto.code());
        }
        if(dto.value() != null && !dto.value().isBlank()) {
            entity.setValue(dto.value());
        }
        if(dto.description() != null && !dto.description().isBlank()) {
            entity.setDescription(dto.description());
        }
        return entity;
    }

    @Override
    public DialectResponse toResponse(Dialect entity) {
        return new DialectResponse(
                entity.getId(),
                entity.getCode(),
                entity.getValue(),
                entity.getDescription(),
                entity.getCreatedAt(),
                entity.getUpdatedAt()
        );
    }
}
