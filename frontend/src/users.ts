import m from 'mithril';
import type { User } from './types/user.ts';

interface UsersResponse {
  users: User[]
}

export function getUsers() {
  return m.request<UsersResponse>({
    method: "GET",
    url: "/api/users"
  });
}
