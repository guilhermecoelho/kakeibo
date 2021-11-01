import { Group } from "./group.model";

export interface Category{
    id: number,
    name: string,
    groupId: number,
    group: Group
  }