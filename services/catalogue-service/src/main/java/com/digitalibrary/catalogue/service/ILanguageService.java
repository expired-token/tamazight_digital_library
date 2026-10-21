package com.digitalibrary.catalogue.service;

import com.digitalibrary.catalogue.dto.language.CreateLanguageRequest;
import com.digitalibrary.catalogue.dto.language.LanguageResponse;
import com.digitalibrary.catalogue.dto.language.UpdateLanguageRequest;
import com.digitalibrary.catalogue.entity.Language;

public interface ILanguageService extends IBaseService<Language, CreateLanguageRequest, UpdateLanguageRequest, LanguageResponse> {}
