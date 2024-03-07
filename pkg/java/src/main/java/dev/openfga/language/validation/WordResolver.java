package dev.openfga.language.validation;

interface WordResolver {
    int resolve(int wordIndex, String rawLine, String symbol);
}
