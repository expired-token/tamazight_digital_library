package com.digitalibrary.catalogue.dto.language;

import java.time.Instant;

public record LanguageResponse(
        Long id,
        String code,
        String value,
        String description,
        Instant createdAt,
        Instant updatedAt
) {}