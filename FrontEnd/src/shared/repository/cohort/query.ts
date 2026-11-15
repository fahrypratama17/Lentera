"use client";

import {
  useMutation,
  useQuery,
  useQueryClient,
  type QueryClient,
} from "@tanstack/react-query";

import {
  cancelInvitationAction,
  createCohortAction,
  deleteCohortAction,
  getCohortDetailAction,
  getCohortInvitationsAction,
  getCohortMembersAction,
  getCohortsAction,
  inviteMemberAction,
  leaveCohortAction,
  removeMemberAction,
  updateCohortAction,
  updateMemberRoleAction,
} from "./action";
import type {
  CohortDetailDto,
  CohortDetailResponseDto,
  CohortDto,
  CohortListResponseDto,
  CohortMutationResponseDto,
  CohortRole,
  CreateCohortDto,
  UpdateCohortDto,
  UploadCohortPdfResponseDto,
} from "./dto";

type Id = string | number | undefined;

export const cohortKeys = {
  all: ["cohorts"] as const,
  list: () => [...cohortKeys.all, "list"] as const,
  detail: (cohortId: Id) => [...cohortKeys.all, "detail", cohortId] as const,
  members: (cohortId: Id) => [...cohortKeys.all, "members", cohortId] as const,
  invitations: (cohortId: Id) =>
    [...cohortKeys.all, "invitations", cohortId] as const,
};

/** Dipakai bersama oleh semua mutation agar cache tetap sinkron. */
function invalidateCohort(client: QueryClient, cohortId?: Id) {
  void client.invalidateQueries({ queryKey: cohortKeys.list() });

  if (cohortId !== undefined) {
    void client.invalidateQueries({ queryKey: cohortKeys.detail(cohortId) });
    void client.invalidateQueries({ queryKey: cohortKeys.members(cohortId) });
    void client.invalidateQueries({
      queryKey: cohortKeys.invitations(cohortId),
    });
  }
}

/* ────────────────────────────── Query ────────────────────────────── */

/** GET /cohorts */
export function useCohortsQuery() {
  return useQuery<CohortListResponseDto, Error, CohortDto[]>({
    queryKey: cohortKeys.list(),
    queryFn: getCohortsAction,
    select: (res) => res.data ?? [],
  });
}

/** GET /cohorts/:id */
export function useCohortDetailQuery(cohortId: Id) {
  return useQuery<CohortDetailResponseDto, Error, CohortDetailDto | null>({
    queryKey: cohortKeys.detail(cohortId),
    queryFn: () => getCohortDetailAction(cohortId!),
    select: (res) => res.data ?? null,
    enabled: !!cohortId,
  });
}

/** GET /cohorts/:id/members */
export function useCohortMembersQuery(cohortId: Id) {
  return useQuery({
    queryKey: cohortKeys.members(cohortId),
    queryFn: () => getCohortMembersAction(cohortId!),
    enabled: !!cohortId,
  });
}

/** GET /cohorts/:id/invitations */
export function useCohortInvitationsQuery(cohortId: Id) {
  return useQuery({
    queryKey: cohortKeys.invitations(cohortId),
    queryFn: () => getCohortInvitationsAction(cohortId!),
    enabled: !!cohortId,
  });
}

/* ────────────────────────────── Mutation ────────────────────────────── */

/** POST /cohorts */
export function useCreateCohortMutation() {
  const client = useQueryClient();

  return useMutation<CohortMutationResponseDto, Error, CreateCohortDto>({
    mutationKey: ["cohorts", "create"],
    mutationFn: createCohortAction,
    onSuccess: () => invalidateCohort(client),
  });
}

/** PUT /cohorts/:id */
export function useUpdateCohortMutation() {
  const client = useQueryClient();

  return useMutation<
    CohortMutationResponseDto,
    Error,
    { cohortId: Id; payload: UpdateCohortDto }
  >({
    mutationKey: ["cohorts", "update"],
    mutationFn: ({ cohortId, payload }) =>
      updateCohortAction(cohortId!, payload),
    onSuccess: (_res, { cohortId }) => invalidateCohort(client, cohortId),
  });
}

/** DELETE /cohorts/:id */
export function useDeleteCohortMutation() {
  const client = useQueryClient();

  return useMutation<CohortMutationResponseDto, Error, Id>({
    mutationKey: ["cohorts", "delete"],
    mutationFn: (cohortId) => deleteCohortAction(cohortId!),
    onSuccess: (_res, cohortId) => invalidateCohort(client, cohortId),
  });
}


export function useUpdateMemberRoleMutation() {
  const client = useQueryClient();

  return useMutation<
    unknown,
    Error,
    { cohortId: Id; userId: Id; role: CohortRole }
  >({
    mutationKey: ["cohorts", "memberRole"],
    mutationFn: ({ cohortId, userId, role }) =>
      updateMemberRoleAction(cohortId!, userId!, role),
    onSuccess: (_res, { cohortId }) => invalidateCohort(client, cohortId),
  });
}

/** DELETE /cohorts/:id/members/:userId */
export function useRemoveMemberMutation() {
  const client = useQueryClient();

  return useMutation<unknown, Error, { cohortId: Id; userId: Id }>({
    mutationKey: ["cohorts", "removeMember"],
    mutationFn: ({ cohortId, userId }) =>
      removeMemberAction(cohortId!, userId!),
    onSuccess: (_res, { cohortId }) => invalidateCohort(client, cohortId),
  });
}

/** POST /cohorts/:id/leave */
export function useLeaveCohortMutation() {
  const client = useQueryClient();

  return useMutation<unknown, Error, Id>({
    mutationKey: ["cohorts", "leave"],
    mutationFn: (cohortId) => leaveCohortAction(cohortId!),
    onSuccess: () => invalidateCohort(client),
  });
}

/** POST /cohorts/:id/invitations */
export function useInviteMemberMutation() {
  const client = useQueryClient();

  return useMutation<
    unknown,
    Error,
    { cohortId: Id; email: string; role: CohortRole }
  >({
    mutationKey: ["cohorts", "invite"],
    mutationFn: ({ cohortId, email, role }) =>
      inviteMemberAction(cohortId!, { email, role }),
    onSuccess: (_res, { cohortId }) => invalidateCohort(client, cohortId),
  });
}

/** DELETE /cohorts/:id/invitations/:invitationId */
export function useCancelInvitationMutation() {
  const client = useQueryClient();

  return useMutation<
    unknown,
    Error,
    { cohortId: Id; invitationId: Id }
  >({
    mutationKey: ["cohorts", "cancelInvitation"],
    mutationFn: ({ cohortId, invitationId }) =>
      cancelInvitationAction(cohortId!, invitationId!),
    onSuccess: (_res, { cohortId }) => invalidateCohort(client, cohortId),
  });
}
