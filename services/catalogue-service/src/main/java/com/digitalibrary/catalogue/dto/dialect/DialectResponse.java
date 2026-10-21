package com.digitalibrary.catalogue.dto.dialect;

import java.time.Instant;

public record DialectResponse(
        Long id,
        String code,
        String value,
        String description,
        Instant createdAt,
        Instant updatedAt
) {}
