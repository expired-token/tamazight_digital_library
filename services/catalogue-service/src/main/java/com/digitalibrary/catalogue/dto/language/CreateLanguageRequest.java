package com.digitalibrary.catalogue.dto.language;

import jakarta.validation.constraints.NotBlank;

public record CreateLanguageRequest(
        @NotBlank(message = "Code is required")
        String code,

        @NotBlank(message = "Value is required")
        String value,

        String description
) {}
