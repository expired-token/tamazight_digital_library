package com.digitalibrary.catalogue.dto.dialect;

public record UpdateDialectRequest(
    String code,
    String value,
    String description
) {}
