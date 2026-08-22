import type { TodoStatus } from "../types/todo";
import { atom } from "jotai";

export type TodoFilterStatus = "all" | TodoStatus;

export const todoFilterAtom = atom<TodoFilterStatus>("all");
