import { api } from "@/shared/lib/axios";

import type { ProfileDto, ProfileResponseDto } from "./dto";

export async function getMyProfileAction(): Promise<ProfileDto> {
  const { data } = await api.get<ProfileResponseDto>("/me");

  if (!data.success || !data.data) {
    throw new Error(data.message ?? "Gagal mengambil data profil.");
  }

  return data.data;
}
