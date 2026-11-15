"use client";

import { useQuery } from "@tanstack/react-query";

import { getMyProfileAction } from "./action";

export function useMyProfileQuery() {
  return useQuery({
    queryKey: ["auth", "me"],
    queryFn: getMyProfileAction,
    staleTime: 60 * 1000,
    retry: false,
  });
}
