package com.digitalibrary.catalogue.service;

import com.digitalibrary.catalogue.dto.dialect.CreateDialectRequest;
import com.digitalibrary.catalogue.dto.dialect.DialectResponse;
import com.digitalibrary.catalogue.dto.dialect.UpdateDialectRequest;
import com.digitalibrary.catalogue.entity.Dialect;

public interface IDialectService extends IBaseService<Dialect, CreateDialectRequest, UpdateDialectRequest, DialectResponse> {}
