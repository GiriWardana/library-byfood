"use client"
// frontend/app/context/BookContext.tsx
import React, { createContext, useState, useContext, ReactNode } from 'react';

interface Book {
  id: number;
  title: string;
  author: string;
  year: number;
}

interface BookProviderProps {
  children: ReactNode;
}

interface BookContextValue {
  books: Book[];
  setBooks: React.Dispatch<React.SetStateAction<Book[]>>;
}

const BookContext = createContext<BookContextValue | undefined>(undefined);

export const BookProvider: React.FC<BookProviderProps> = ({ children }) => {
  const [books, setBooks] = useState<Book[]>([]);
  return (
    <BookContext.Provider value={{ books, setBooks }}>
      {children}
    </BookContext.Provider>
  );
};

export const useBooks = () => {
  const context = useContext(BookContext);
  if (!context) {
    throw new Error('useBooks must be used within a BookProvider');
  }
  return context;
};