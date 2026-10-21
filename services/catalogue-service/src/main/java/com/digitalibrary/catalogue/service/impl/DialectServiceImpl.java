package com.digitalibrary.catalogue.service.impl;

import com.digitalibrary.catalogue.dto.dialect.CreateDialectRequest;
import com.digitalibrary.catalogue.dto.dialect.DialectResponse;
import com.digitalibrary.catalogue.dto.dialect.UpdateDialectRequest;
import com.digitalibrary.catalogue.entity.Dialect;
import com.digitalibrary.catalogue.exception.RecordNotFoundException;
import com.digitalibrary.catalogue.exception.RequiredRequestParameterMissingException;
import com.digitalibrary.catalogue.mapper.DialectMapper;
import com.digitalibrary.catalogue.repository.IDialectRepository;
import com.digitalibrary.catalogue.service.IDialectService;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class DialectServiceImpl implements IDialectService {

    private final IDialectRepository repository;
    private final DialectMapper mapper;

    public DialectServiceImpl(IDialectRepository dialectRepository, DialectMapper dialectMapper) {
        repository = dialectRepository;
        mapper = dialectMapper;
    }

    @Override
    public List<DialectResponse> searchAll() {
        List<Dialect> dialectList = repository.findAll();
        return mapper.toResponseAll(dialectList);
    }

    @Override
    public DialectResponse searchByCode(String code) {
        if(code.isBlank()) {
            throw new RequiredRequestParameterMissingException("A required parameter is missing");
        }

        Dialect entity = repository.findByCode(code).orElseThrow(() -> new RecordNotFoundException("Dialect with code '" + code + "' is not found"));

        return mapper.toResponse(entity);
    }

    @Override
    public DialectResponse create(CreateDialectRequest dto) {
        if(dto.code().isBlank() || dto.value().isBlank()) {
            throw new RequiredRequestParameterMissingException("Required parameter(s) are missing");
        }

        Dialect entity = repository.save(mapper.createRequestToEntity(dto));

        return mapper.toResponse(entity);
    }

    @Override
    public DialectResponse update(Long id, UpdateDialectRequest dto) {
        Dialect entity = repository.findById(id).orElseThrow(() -> new RecordNotFoundException("Dialect not found"));
        Dialect updated = repository.save(mapper.updateRequestToEntity(entity, dto));

        return mapper.toResponse(updated);
    }

    @Override
    public void delete(Long id) {
        if(!repository.existsById(id)) {
            throw new RecordNotFoundException("Dialect not found");
        }

        repository.deleteById(id);
    }
}
