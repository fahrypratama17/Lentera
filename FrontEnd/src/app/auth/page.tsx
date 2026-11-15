import type { Metadata } from "next";

import Auth from "@/features/auth/container/Auth";

export const metadata: Metadata = {
  title: "Login & Register",
  description: "Masuk atau daftar untuk mulai belajar",
};

export default function Page() {
  return <Auth />;
}
