"use client";

import { useEffect, useState } from "react";
import type { FormEvent } from "react";
import { isAxiosError } from "axios";

import { useMyProfileQuery } from "@/shared/repository/profile/query";
import type { ProfileErrorDto } from "@/shared/repository/profile/dto";

function getErrorMessage(error: unknown): string {
  if (isAxiosError<ProfileErrorDto>(error)) {
    if (error.response?.status === 401) {
      return "Sesi login tidak valid atau telah berakhir. Silakan login kembali.";
    }

    if (!error.response) {
      return "Gagal terhubung ke server.";
    }

    return (
      error.response.data?.message ??
      error.response.data?.error ??
      "Terjadi kesalahan saat mengambil profil."
    );
  }

  return error instanceof Error
    ? error.message
    : "Terjadi kesalahan yang tidak diketahui.";
}

export function useProfile() {
  const profileQuery = useMyProfileQuery();

  const user = profileQuery.data;

  const [localName, setLocalName] = useState<string | null>(null);
  const [name, setName] = useState("");

  const [nameError, setNameError] = useState("");
  const [nameSuccess, setNameSuccess] = useState("");

  const [currentPassword, setCurrentPassword] = useState("");
  const [newPassword, setNewPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");

  const [passwordError, setPasswordError] = useState("");
  const [passwordNotice, setPasswordNotice] = useState("");

  useEffect(() => {
    if (!user) return;

    const storageKey = `studytrack-profile-name-${user.id}`;

    let storedName: string | null = null;

    try {
      storedName = localStorage.getItem(storageKey);
    } catch {

    }

    setLocalName(storedName);
    setName(storedName ?? user.name);
  }, [user]);

  const currentName = localName ?? user?.name ?? "";
  const email = user?.email ?? "";

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

    if (!user) {
      setNameError("Data pengguna belum tersedia.");
      return;
    }

    const trimmedName = name.trim();

    if (trimmedName.length < 2 || trimmedName.length > 100) {
      setNameError("Nama harus terdiri dari 2–100 karakter.");
      return;
    }

    if (trimmedName === currentName) return;

    try {
      localStorage.setItem(`studytrack-profile-name-${user.id}`, trimmedName);

      setLocalName(trimmedName);
      setName(trimmedName);

      setNameSuccess(
        "Nama tampilan disimpan di browser. Nama akun di server belum berubah.",
      );
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

    // TODO: Integrasi API perubahan password.

    setCurrentPassword("");
    setNewPassword("");
    setConfirmPassword("");

    setPasswordNotice(
      "Validasi berhasil. Password belum diubah karena belum terhubung ke backend.",
    );
  }

  return {
    ready: !profileQuery.isPending,
    loading: profileQuery.isPending,
    loadError: profileQuery.isError ? getErrorMessage(profileQuery.error) : "",
    reloadProfile: profileQuery.refetch,

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
