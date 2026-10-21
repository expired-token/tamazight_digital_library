package com.digitalibrary.catalogue.mapper;

import com.digitalibrary.catalogue.entity.BaseEntity;

import java.util.ArrayList;
import java.util.List;

public abstract class BaseMapper<E extends BaseEntity, CR, UR, R> {
    public abstract E createRequestToEntity(CR dto);
    public abstract E updateRequestToEntity(E entity, UR dto);
    public abstract R toResponse(E entity);

    public List<R> toResponseAll(List<E> entityList) {
        List<R> responseList = new ArrayList<>();

        for(E entity: entityList) {
            responseList.add(toResponse(entity));
        }

        return responseList;
    }
}
