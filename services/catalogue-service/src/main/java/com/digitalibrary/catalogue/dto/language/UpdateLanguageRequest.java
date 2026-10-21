package com.digitalibrary.catalogue.dto.language;

public record UpdateLanguageRequest(
    String code,
    String value,
    String description
) {}
