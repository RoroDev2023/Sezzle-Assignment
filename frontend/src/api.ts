export type Operation =
  | 'add'
  | 'subtract'
  | 'multiply'
  | 'divide'
  | 'power'
  | 'sqrt'
  | 'percentage';

export type CalculatePayload = {
  operation: Operation;
  a: number;
  b?: number;
};

export type CalculateResponse = {
  operation: Operation;
  result: number;
};

type ApiErrorResponse = {
  error: string;
};

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL ?? 'http://localhost:8080';

export async function calculate(payload: CalculatePayload): Promise<CalculateResponse> {
  const response = await fetch(`${API_BASE_URL}/api/calculate`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(payload),
  });

  const data = (await response.json()) as CalculateResponse | ApiErrorResponse;

  if (!response.ok) {
    throw new Error('error' in data ? data.error : 'Calculation failed');
  }

  return data as CalculateResponse;
}

