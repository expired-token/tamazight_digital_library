package com.digitalibrary.catalogue.service;

import com.digitalibrary.catalogue.entity.BaseEntity;

import java.util.List;

public interface IBaseService<E extends BaseEntity, CR, UR, R> {
    List<R> searchAll();
    R searchByCode(String code);
    R create(CR dto);
    R update(Long id, UR dto);
    void delete(Long id);
}
