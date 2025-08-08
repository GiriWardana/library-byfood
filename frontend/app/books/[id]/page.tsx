"use client";

import { useEffect, useState } from "react";
import { useParams, useRouter } from "next/navigation";

interface Book {
  id: number;
  title: string;
  author: string;
  year: number;
}

const BookDetailPage = () => {
  const { id } = useParams();
  const router = useRouter();
  const [book, setBook] = useState<Book | null>(null);
  useEffect(() => {
    fetch(`${process.env.NEXT_PUBLIC_API_URL}/books/${id}`)
      .then((res) => res.json())
      .then((data) => setBook(data))
      .catch((err) => console.error("Error fetching book:", err));
  }, [id]);

  if (!book) return <p className="text-center mt-10">Loading...</p>;

  return (
    <div className="max-w-2xl mx-auto mt-10 p-6 bg-white shadow rounded">
      <button   
        onClick={() => router.push('/')}
        className="bg-gray-200 hover:bg-gray-300 text-black py-2 px-4 rounded"
      >
        ← Back to Home
      </button>
      <h1 className="text-2xl font-bold mb-4">{book.title}</h1>
      <p className="text-gray-700"><strong>Author:</strong> {book.author}</p>
      <p className="text-gray-700"><strong>Year:</strong> {book.year}</p>
    </div>
  );
};

export default BookDetailPage;
