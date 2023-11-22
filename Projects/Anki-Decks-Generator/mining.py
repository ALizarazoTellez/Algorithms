#!/usr/bin/env python3

"""Mining

This command generates CSV files containing Anki decks, for apply the method
Mining I+1.

TODO: Translate to English (it is in Spanish).

El primer argumento recibe un archivo con frases, una en cada línea.
El segundo argumento es un archivo donde escribir las frases a memorizar, el
contenido anterior se elimina.

El criterio para añadir una frase a memorizar es que solo desconozcas una
palabra.
"""

import sys

KNOW_WORDS_FILE = './know-words.txt'


def main(*args) -> int:
    """Main routine."""

    if len(args) != 3:
        print("Args: phrases_source mining_phrases")
        return 1

    know_words = load_words(KNOW_WORDS_FILE)
    phrases = load_phrases(args[1])

    learn_phrases = []

    for phrase in phrases:
        words = unknown_words(phrase, know_words)

        if len(words) == 0:
            continue

        if len(words) == 1:
            learn_phrases.append(phrase)
            continue

        # At this point, there are two or more possibly unknown words.
        new_unknown_words = check_unknown_words(words)

        if len(new_unknown_words) == 1:
            learn_phrases.append(phrase)

        know_words.update(words - new_unknown_words)

    save_words(KNOW_WORDS_FILE, know_words)
    save_phrases(args[2], learn_phrases)

    return 0


def load_words(path: str) -> set[str]:
    lines: list[str]
    with open(path) as file:
        lines = file.readlines()

    words = set()
    for line in lines:
        words.add(line.strip())

    return words


def save_words(path: str, words: set[str]):
    with open(path, 'w') as file:
        for word in words:
            file.write(word+'\n')


def load_phrases(path: str) -> set[str]:
    lines: list[str]
    with open(path) as file:
        lines = file.readlines()

    phrases = set()
    for line in lines:
        phrases.add(line.strip())

    return phrases


def save_phrases(path: str, phrases: set[str]):
    with open(path, 'w') as file:
        for phrase in phrases:
            file.write(phrase + '\n')


def unknown_words(phrase: str, words: tuple[str]) -> set[str]:
    unknown = set()

    phrase_words = phrase.split()

    for word in phrase_words:
        if word not in words:
            unknown.add(simplify_word(word))

    return unknown


def check_unknown_words(words: tuple[str]) -> set[str]:
    unknown = set()

    for word in words:
        decision = input(f'Do you know the word "{word}"? [y/N]: ')

        if decision.lower() == 'y':
            continue

        unknown.add(simplify_word(word))

    return unknown


def simplify_word(word: str) -> str:
    """Return a word without punctuation signs."""

    PUNCTUATION = {
        '.',
        ',',
        '"',
        # The character `'` is used for contractions in the English language.
        # "'",
        '¿', '?',
        '¡', '!',
    }

    for punctuation in PUNCTUATION:
        word = word.replace(punctuation, '', 1)

    return word


if __name__ == '__main__':
    exit(main(*sys.argv))
