import type { Metadata } from "next";

import Profile from "@/features/profile/container/Profile";

export const metadata: Metadata = {
  title: "Profil Saya | StudyTrack",
  description: "Kelola informasi pribadi dan keamanan akun.",
};

export default function Page() {
  return <Profile />;
}
