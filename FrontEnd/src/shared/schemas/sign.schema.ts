import { z } from "zod";
import type { SignMode } from "@/features/auth/types/type";

export function createSignSchema(mode: SignMode) {
  return z
    .object({
      userName: z.string(),

      email: z
        .string()
        .trim()
        .min(1, "Email wajib diisi")
        .email("Format email tidak valid"),

      password: z.string().min(1, "Password wajib diisi"),

      confirmPassword: z.string(),
    })
    .superRefine((data, ctx) => {
      if (mode === "signIn") return;

      if (data.userName.trim().length < 3) {
        ctx.addIssue({
          code: "custom",
          path: ["userName"],
          message: "Username minimal 3 karakter",
        });
      }

      if (data.password.length < 8) {
        ctx.addIssue({
          code: "custom",
          path: ["password"],
          message: "Password minimal 8 karakter",
        });
      }

      if (data.confirmPassword !== data.password) {
        ctx.addIssue({
          code: "custom",
          path: ["confirmPassword"],
          message: "Konfirmasi password tidak cocok",
        });
      }
    });
}
