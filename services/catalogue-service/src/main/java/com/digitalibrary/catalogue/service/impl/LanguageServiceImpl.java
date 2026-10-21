package com.digitalibrary.catalogue.service.impl;

import com.digitalibrary.catalogue.dto.language.CreateLanguageRequest;
import com.digitalibrary.catalogue.dto.language.LanguageResponse;
import com.digitalibrary.catalogue.dto.language.UpdateLanguageRequest;
import com.digitalibrary.catalogue.entity.Language;
import com.digitalibrary.catalogue.exception.RecordNotFoundException;
import com.digitalibrary.catalogue.exception.RequiredRequestParameterMissingException;
import com.digitalibrary.catalogue.mapper.LanguageMapper;
import com.digitalibrary.catalogue.repository.ILanguageRepository;
import com.digitalibrary.catalogue.service.ILanguageService;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class LanguageServiceImpl implements ILanguageService {

    private final ILanguageRepository repository;
    private final LanguageMapper mapper;

    public LanguageServiceImpl(ILanguageRepository languageRepository, LanguageMapper languageMapper) {
        repository = languageRepository;
        mapper = languageMapper;
    }


    @Override
    public List<LanguageResponse> searchAll() {
        List<Language> languageList = repository.findAll();
        return mapper.toResponseAll(languageList);
    }

    @Override
    public LanguageResponse searchByCode(String code) {
        if(code.isBlank()) {
            throw new RequiredRequestParameterMissingException("A required parameter is missing");
        }

        Language entity = repository.findByCode(code).orElseThrow(() -> new RecordNotFoundException("Language with code '" + code + "' is not found"));

        return mapper.toResponse(entity);
    }

    @Override
    public LanguageResponse create(CreateLanguageRequest dto) {
        if(dto.code().isBlank() || dto.value().isBlank()) {
            throw new RequiredRequestParameterMissingException("Required parameter(s) are missing");
        }

        Language entity = repository.save(mapper.createRequestToEntity(dto));

        return mapper.toResponse(entity);
    }

    @Override
    public LanguageResponse update(Long id, UpdateLanguageRequest dto) {
        Language entity = repository.findById(id).orElseThrow(() -> new RequiredRequestParameterMissingException("Required parameter is missing"));
        Language updated = repository.save(mapper.updateRequestToEntity(entity, dto));

        return mapper.toResponse(updated);
    }

    @Override
    public void delete(Long id) {
        if(!repository.existsById(id)) {
            throw new RequiredRequestParameterMissingException("Required parameter is missing");
        }

        repository.deleteById(id);
    }
}
