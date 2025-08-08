"use client";

import React, { useEffect, useState } from "react";
import { useBooks } from "./context/BookContext";
import { useToast } from "./components/ToastProvider";
import Link from "next/link";
// import { NextConfig } from "next";
import nextConfig from "@/next.config";

interface Book {
  id: number;
  title: string;
  author: string;
  year: number;
}

const HomePage: React.FC = () => {
  const { showToast } = useToast();
  const { books, setBooks } = useBooks();
  const [newBook, setNewBook] = useState<Omit<Book, "id">>({
    title: "",
    author: "",
    year: 0,
  });
  const apiUrl = nextConfig?.env?.NEXT_PUBLIC_API_URL;
  console.log("🚀 ~ HomePage ~ apiUrl:", apiUrl)

  const [editBook, setEditBook] = useState<Book | null>(null);
  const [deleteTarget, setDeleteTarget] = useState<Book | null>(null);
  const [deleteConfirmInput, setDeleteConfirmInput] = useState("");

  useEffect(() => {
    fetch(`${apiUrl}/books`)
      .then((res) => res.json())
      .then((data) => setBooks(data))
      .catch((err) => console.error("Error fetching books:", err));
  }, [setBooks]);

  const handleCreate = () => {
    if (!newBook.title || !newBook.author || !newBook.year) {
      showToast("All fields are required.", "error");
      return;
    }
    fetch(`${process.env.NEXT_PUBLIC_API_URL}/books`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(newBook),
    })
      .then(async (res) => {
        const data = await res.json();
        if (!res.ok) {
          showToast(data.error || "Failed to create book", "error");
          throw new Error(data.error || "Request failed");
        }
        setBooks((prev) => [...prev, data]);
        setNewBook({ title: "", author: "", year: 0 });
        showToast("Book created successfully!", "success");
      })
      .catch((err) => {
        console.error("Error creating book:", err);
      });
  };

  const handleEditSubmit = () => {
    if (!editBook) return;
    fetch(`${process.env.NEXT_PUBLIC_API_URL}/books/${editBook.id}`, {
      method: "PUT",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(editBook),
    })
     .then(async (res) => {
        const data = await res.json();
        if (!res.ok) {
          showToast(data.error || "Failed to update book", "error");
          throw new Error(data.error);
        }
        setBooks((prev) => prev.map((book) => (book.id === data.id ? data : book)));
        setEditBook(null);
        showToast("Book updated successfully!", "success");
      })

  };

  const handleDelete = () => {
    if (!deleteTarget || deleteConfirmInput !== deleteTarget.title) return;

    fetch(`${process.env.NEXT_PUBLIC_API_URL}/books/${deleteTarget.id}`, {
      method: "DELETE",
    })
      .then((res) => {
        if (!res.ok) {
          showToast("Failed to delete book", "error");
          throw new Error("Delete failed");
        }
        setBooks((prev) => prev.filter((book) => book.id !== deleteTarget.id));
        setDeleteTarget(null);
        setDeleteConfirmInput("");
        showToast("Book deleted successfully!", "success");
      })
  };

  return (
    <div className="max-w-4xl mx-auto p-6">
      <h1 className="text-3xl font-bold mb-6 text-center text-gray-800">📚 Book Manager</h1>

      {/* Add Book Form */}
      <div className="bg-white p-6 rounded-lg shadow mb-8">
        <h2 className="text-xl font-semibold mb-4 text-gray-700">Add New Book</h2>
        <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
          <input
            type="text"
            placeholder="Title"
            value={newBook.title}
            onChange={(e) => setNewBook({ ...newBook, title: e.target.value })}
            className="border border-gray-300 rounded px-4 py-2"
          />
          <input
            type="text"
            placeholder="Author"
            value={newBook.author}
            onChange={(e) => setNewBook({ ...newBook, author: e.target.value })}
            className="border border-gray-300 rounded px-4 py-2"
          />
          <input
            type="number"
            placeholder="Year"
            value={newBook.year}
            onChange={(e) => setNewBook({ ...newBook, year: Number(e.target.value) })}
            className="border border-gray-300 rounded px-4 py-2"
          />
        </div>
        <button
          onClick={handleCreate}
          className="mt-4 bg-blue-600 hover:bg-blue-700 text-white font-semibold py-2 px-6 rounded"
        >
          Add Book
        </button>
      </div>

      {/* Book List */}
      <div>
        <h2 className="text-2xl font-semibold mb-4 text-gray-800">Book List</h2>
        {books.length === 0 ? (
          <p className="text-gray-500">No books found.</p>
        ) : (
          <ul className="space-y-4">
            {books.map((book) => (
              <li
                key={book.id}
                className="bg-gray-100 p-4 rounded-lg shadow flex justify-between items-center"
              >
                {/* <div>
                  <h3 className="text-lg font-bold text-gray-700">{book.title}</h3>
                  <p className="text-sm text-gray-600">
                    {book.author} ({book.year})
                  </p>
                </div> */}
                <Link href={`/books/${book.id}`}>
                  <h3 className="text-lg font-bold text-blue-600 hover:underline">{book.title}</h3>
                  <p className="text-sm text-gray-600">
                    {book.author} ({book.year})
                  </p>
                </Link>

                <div className="flex gap-2">
                  <button
                    onClick={() => setEditBook(book)}
                    className="bg-yellow-500 hover:bg-yellow-600 text-white px-3 py-1 rounded text-sm"
                  >
                    Edit
                  </button>
                  <button
                    onClick={() => setDeleteTarget(book)}
                    className="bg-red-500 hover:bg-red-600 text-white px-3 py-1 rounded text-sm"
                  >
                    Delete
                  </button>
                </div>
              </li>
            ))}
          </ul>
        )}
      </div>

      {/* Edit Modal */}
      {editBook && (
        <div className="fixed inset-0 bg-black bg-opacity-40 flex items-center justify-center z-50">
          <div className="bg-white rounded-lg p-6 w-full max-w-md">
            <h2 className="text-xl font-bold mb-4 text-gray-700">Edit Book</h2>
            <input
              type="text"
              value={editBook.title}
              onChange={(e) => setEditBook({ ...editBook, title: e.target.value })}
              className="mb-3 w-full border border-gray-300 rounded px-4 py-2"
            />
            <input
              type="text"
              value={editBook.author}
              onChange={(e) => setEditBook({ ...editBook, author: e.target.value })}
              className="mb-3 w-full border border-gray-300 rounded px-4 py-2"
            />
            <input
              type="number"
              value={editBook.year}
              onChange={(e) => setEditBook({ ...editBook, year: Number(e.target.value) })}
              className="mb-4 w-full border border-gray-300 rounded px-4 py-2"
            />
            <div className="flex justify-end gap-3">
              <button
                onClick={() => setEditBook(null)}
                className="px-4 py-2 rounded bg-gray-300 hover:bg-gray-400"
              >
                Cancel
              </button>
              <button
                onClick={handleEditSubmit}
                className="px-4 py-2 rounded bg-green-600 hover:bg-green-700 text-white"
              >
                Save
              </button>
            </div>
          </div>
        </div>
      )}

      {/* Delete Confirmation Modal */}
      {deleteTarget && (
        <div className="fixed inset-0 bg-black bg-opacity-40 flex items-center justify-center z-50">
          <div className="bg-white rounded-lg p-6 w-full max-w-md">
            <h2 className="text-xl font-bold text-red-600 mb-4">Confirm Deletion</h2>
            <p className="mb-3">
              Type <strong>{deleteTarget.title}</strong> to confirm deletion:
            </p>
            <input
              type="text"
              value={deleteConfirmInput}
              onChange={(e) => setDeleteConfirmInput(e.target.value)}
              className="w-full border border-gray-300 rounded px-4 py-2 mb-4"
            />
            <div className="flex justify-end gap-3">
              <button
                onClick={() => {
                  setDeleteTarget(null);
                  setDeleteConfirmInput("");
                }}
                className="px-4 py-2 rounded bg-gray-300 hover:bg-gray-400"
              >
                Cancel
              </button>
              <button
                onClick={handleDelete}
                disabled={deleteConfirmInput !== deleteTarget.title}
                className={`px-4 py-2 rounded text-white ${
                  deleteConfirmInput === deleteTarget.title
                    ? "bg-red-600 hover:bg-red-700"
                    : "bg-red-300 cursor-not-allowed"
                }`}
              >
                Delete
              </button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
};

export default HomePage;
