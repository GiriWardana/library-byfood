// frontend/app/layout.tsx
"use client";

import React, { ReactNode } from 'react';
import { BookProvider } from './context/BookContext';
import { ToastProvider } from './components/ToastProvider';
import './globals.css';

interface LayoutProps {
  children: ReactNode;
}

const Layout: React.FC<LayoutProps> = ({ children }) => {
  return (
    <html lang="en">
      <head>
        <title>Library Application</title>
      </head>
      <body>
        <BookProvider>
          <header>
            <h1>Library Application</h1>
          </header>
          <main>
            <ToastProvider>
                {children}
            </ToastProvider>
          </main>
        </BookProvider>
      </body>
    </html>
  );
};

export default Layout;