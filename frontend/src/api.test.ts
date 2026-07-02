import { afterEach, describe, expect, it, vi } from 'vitest';
import { calculate } from './api';

describe('calculate', () => {
  afterEach(() => {
    vi.restoreAllMocks();
  });

  it('returns the API calculation result', async () => {
    vi.spyOn(globalThis, 'fetch').mockResolvedValue(
      new Response(JSON.stringify({ operation: 'multiply', result: 21 }), {
        status: 200,
        headers: { 'Content-Type': 'application/json' },
      }),
    );

    await expect(calculate({ operation: 'multiply', a: 3, b: 7 })).resolves.toEqual({
      operation: 'multiply',
      result: 21,
    });
  });

  it('throws the API error message', async () => {
    vi.spyOn(globalThis, 'fetch').mockResolvedValue(
      new Response(JSON.stringify({ error: 'division by zero' }), {
        status: 400,
        headers: { 'Content-Type': 'application/json' },
      }),
    );

    await expect(calculate({ operation: 'divide', a: 3, b: 0 })).rejects.toThrow('division by zero');
  });
});

