import type { Cohort } from "@/features/dashboard/types/type";

export const mockCohorts: Cohort[] = [
  {
    id: "frontend-development",
    title: "Frontend Development",
    description: "Belajar React, Next.js, dan Tailwind CSS.",
    mentorName: "Atha Tsaqif",
    period: "November - Desember 2026",
    role: "mentor",
    memberCount: 24,
    materialCount: 8,
    previewQuestion:
      "Buatlah landing page responsif menggunakan Next.js dan Tailwind CSS.",
    tone: "ocean",
  },
  {
    id: "backend-engineering",
    title: "Backend Engineering",
    description: "REST API, database, dan autentikasi.",
    mentorName: "Nadia Putri",
    period: "November - Desember 2026",
    role: "member",
    memberCount: 18,
    materialCount: 6,
    previewQuestion:
      "Implementasikan REST API untuk mengelola data pengguna dan autentikasi.",
    tone: "teal",
  },
  {
    id: "ui-ux-design",
    title: "UI/UX Design",
    description: "User research, wireframe, dan prototyping.",
    mentorName: "Raka Pratama",
    period: "Oktober - Desember 2026",
    role: "member",
    memberCount: 16,
    materialCount: 10,
    previewQuestion:
      "Buat wireframe aplikasi mobile berdasarkan studi kasus yang diberikan.",
    tone: "slate",
  },
  {
    id: "devsecops",
    title: "DevSecOps Fundamentals",
    description: "CI/CD, security testing, dan deployment.",
    mentorName: "Atha Tsaqif",
    period: "November 2026",
    role: "mentor",
    memberCount: 20,
    materialCount: 7,
    previewQuestion:
      "Buat pipeline CI/CD yang menjalankan automated security scanning.",
    tone: "teal",
  },
  {
    id: "product-management",
    title: "Product Management",
    description: "Mengembangkan produk berbasis kebutuhan pengguna.",
    mentorName: "Dinda Maharani",
    period: "September - November 2026",
    role: "member",
    memberCount: 22,
    materialCount: 9,
    previewQuestion:
      "Susun Product Requirement Document berdasarkan permasalahan pengguna.",
    tone: "ocean",
  },
  {
    id: "data-science",
    title: "Data Science Basics",
    description: "Python, analisis data, dan visualisasi.",
    mentorName: "Farhan Akbar",
    period: "November - Desember 2026",
    role: "member",
    memberCount: 14,
    materialCount: 5,
    previewQuestion:
      "Lakukan eksplorasi dataset dan buat visualisasi hasil analisis.",
    tone: "slate",
  },
];
