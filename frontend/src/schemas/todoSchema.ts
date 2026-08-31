import { z } from "zod";

export const todoFormSchema = z.object({
  title: z
    .string()
    .trim()
    .max(100, "100文字以内で入力してください")
    .min(1, "タイトルは必須です"),
});

export type TodoFormValues = z.infer<typeof todoFormSchema>;
