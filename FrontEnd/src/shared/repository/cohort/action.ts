import { api } from "@/shared/lib/axios";

import type {
  CohortDetailResponseDto,
  CohortListResponseDto,
  CohortMutationResponseDto,
  CohortRole,
  CreateCohortDto,
  UpdateCohortDto,
  UploadCohortPdfResponseDto,
} from "./dto";

type Id = string | number;

const endpoints = {
  list: () => "/cohorts",
  create: () => "/cohorts",
  detail: (cohortId: Id) => `/cohorts/${cohortId}`,
  update: (cohortId: Id) => `/cohorts/${cohortId}`,
  remove: (cohortId: Id) => `/cohorts/${cohortId}`,
  uploadPdf: (cohortId: Id) => `/cohorts/${cohortId}/pdf`,
  members: (cohortId: Id) => `/cohorts/${cohortId}/members`,
  memberRole: (cohortId: Id, userId: Id) =>
    `/cohorts/${cohortId}/members/${userId}/role`,
  removeMember: (cohortId: Id, userId: Id) =>
    `/cohorts/${cohortId}/members/${userId}`,
  leave: (cohortId: Id) => `/cohorts/${cohortId}/leave`,
  invitations: (cohortId: Id) => `/cohorts/${cohortId}/invitations`,
  cancelInvitation: (cohortId: Id, invitationId: Id) =>
    `/cohorts/${cohortId}/invitations/${invitationId}`,
} as const;


export async function getCohortsAction(): Promise<CohortListResponseDto> {
  const { data } = await api.get<CohortListResponseDto>(endpoints.list());

  return data;
}

export async function createCohortAction(
  payload: CreateCohortDto,
): Promise<CohortMutationResponseDto> {
  const { data } = await api.post<CohortMutationResponseDto>(
    endpoints.create(),
    payload,
  );

  return data;
}

export async function getCohortDetailAction(
  cohortId: Id,
): Promise<CohortDetailResponseDto> {
  const { data } = await api.get<CohortDetailResponseDto>(
    endpoints.detail(cohortId),
  );

  return data;
}

export async function updateCohortAction(
  cohortId: Id,
  payload: UpdateCohortDto,
): Promise<CohortMutationResponseDto> {
  const { data } = await api.put<CohortMutationResponseDto>(
    endpoints.update(cohortId),
    payload,
  );

  return data;
}

export async function deleteCohortAction(
  cohortId: Id,
): Promise<CohortMutationResponseDto> {
  const { data } = await api.delete<CohortMutationResponseDto>(
    endpoints.remove(cohortId),
  );

  return data;
}

/**
 * POST /cohorts/:id/pdf — multipart/form-data.
 * Backend belum mengekspos route ini (pkg/storage masih TODO),
 * jadi path di atas asumsi konvensi. Sesuaikan bila route final berbeda.
 */
export async function uploadCohortPdfAction(
  cohortId: Id,
  file: File,
  documentType?: string,
): Promise<UploadCohortPdfResponseDto> {
  const formData = new FormData();
  formData.append("file", file);
  if (documentType) formData.append("document_type", documentType);

  const { data } = await api.post<UploadCohortPdfResponseDto>(
    endpoints.uploadPdf(cohortId),
    formData,
    { headers: { "Content-Type": "multipart/form-data" } },
  );

  return data;
}


export async function getCohortMembersAction(cohortId: Id) {
  const { data } = await api.get(endpoints.members(cohortId));

  return data;
}

export async function updateMemberRoleAction(
  cohortId: Id,
  userId: Id,
  role: CohortRole,
) {
  const { data } = await api.patch(endpoints.memberRole(cohortId, userId), {
    role,
  });

  return data;
}

export async function removeMemberAction(cohortId: Id, userId: Id) {
  const { data } = await api.delete(endpoints.removeMember(cohortId, userId));

  return data;
}

export async function leaveCohortAction(cohortId: Id) {
  const { data } = await api.post(endpoints.leave(cohortId));

  return data;
}


export async function getCohortInvitationsAction(cohortId: Id) {
  const { data } = await api.get(endpoints.invitations(cohortId));

  return data;
}

export async function inviteMemberAction(
  cohortId: Id,
  payload: { email: string; role: CohortRole },
) {
  const { data } = await api.post(endpoints.invitations(cohortId), payload);

  return data;
}

export async function cancelInvitationAction(
  cohortId: Id,
  invitationId: Id,
) {
  const { data } = await api.delete(
    endpoints.cancelInvitation(cohortId, invitationId),
  );

  return data;
}
