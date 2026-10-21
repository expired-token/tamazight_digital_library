package com.digitalibrary.catalogue.exception;

public class RequiredRequestParameterMissingException extends RuntimeException {
    public RequiredRequestParameterMissingException(String message) {
        super(message);
    }
}
