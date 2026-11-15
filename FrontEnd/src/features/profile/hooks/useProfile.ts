"use client";

import { useEffect, useState } from "react";
import type { FormEvent } from "react";

const STORAGE_KEY = "studytrack-profile-demo";

const demoUser = {
  name: "Pengguna Demo",
  email: "demo@example.com",
};

export function useProfile() {
  const [ready, setReady] = useState(false);

  const [currentName, setCurrentName] = useState(demoUser.name);
  const [name, setName] = useState(demoUser.name);
  const [nameError, setNameError] = useState("");
  const [nameSuccess, setNameSuccess] = useState("");

  const [currentPassword, setCurrentPassword] = useState("");
  const [newPassword, setNewPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");

  const [passwordError, setPasswordError] = useState("");
  const [passwordNotice, setPasswordNotice] = useState("");

  const email = demoUser.email;

  useEffect(() => {
    try {
      const storedName = localStorage.getItem(STORAGE_KEY);

      if (storedName) {
        setCurrentName(storedName);
        setName(storedName);
      }
    } catch {

    } finally {
      setReady(true);
    }
  }, []);

  const initials = currentName
    .trim()
    .split(/\s+/)
    .filter(Boolean)
    .slice(0, 2)
    .map((word) => word.charAt(0).toUpperCase())
    .join("");

  function saveName(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();

    setNameError("");
    setNameSuccess("");

    const trimmedName = name.trim();

    if (trimmedName.length < 2 || trimmedName.length > 100) {
      setNameError("Nama harus terdiri dari 2–100 karakter.");
      return;
    }

    if (trimmedName === currentName) return;

    try {
      localStorage.setItem(STORAGE_KEY, trimmedName);

      setCurrentName(trimmedName);
      setName(trimmedName);

      setNameSuccess("Nama berhasil disimpan di browser (mode demo).");
    } catch {
      setNameError("Gagal menyimpan nama di browser.");
    }
  }

  function changePassword(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();

    setPasswordError("");
    setPasswordNotice("");

    if (!currentPassword) {
      setPasswordError("Password saat ini wajib diisi.");
      return;
    }

    if (newPassword.length < 8) {
      setPasswordError("Password baru minimal 8 karakter.");
      return;
    }

    if (newPassword !== confirmPassword) {
      setPasswordError("Konfirmasi password tidak cocok.");
      return;
    }

    if (currentPassword === newPassword) {
      setPasswordError("Password baru harus berbeda dari password lama.");
      return;
    }

    // TODO: Integrate with backend.

    setCurrentPassword("");
    setNewPassword("");
    setConfirmPassword("");

    setPasswordNotice(
      "Validasi berhasil. Password belum diubah karena fitur ini belum terhubung ke backend.",
    );
  }

  return {
    ready,
    currentName,
    email,
    initials,

    name,
    setName,
    nameError,
    nameSuccess,
    saveName,

    currentPassword,
    setCurrentPassword,
    newPassword,
    setNewPassword,
    confirmPassword,
    setConfirmPassword,
    passwordError,
    passwordNotice,
    changePassword,
  };
}
