"use client";

import { useRouter } from "next/navigation";
import { useQueryClient } from "@tanstack/react-query";
import { isAxiosError } from "axios";
import { useSignOutMutation } from "@/shared/repository/auth/query";
import { removeAccessToken } from "@/shared/lib/auth-storage";

export function useLogout() {
  const router = useRouter();
  const queryClient = useQueryClient();

  const logoutMutation = useSignOutMutation();

  const clearSession = () => {
    removeAccessToken();
    queryClient.clear();

    router.replace("/auth");
    router.refresh();
  };

  const logout = async () => {
    if (logoutMutation.isPending) return;

    try {
      await logoutMutation.mutateAsync();

      clearSession();
    } catch (error) {
      if (isAxiosError(error) && error.response?.status === 401) {
        clearSession();
        return;
      }
      console.error("Logout failed:", error);
    }
  };

  return {
    logout,
    isPending: logoutMutation.isPending,
    error: logoutMutation.isError ? "Gagal logout. Silakan coba lagi." : null,
  };
}
