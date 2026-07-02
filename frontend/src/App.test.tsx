import { render, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { afterEach, describe, expect, it, vi } from 'vitest';
import { App } from './App';

describe('App', () => {
  afterEach(() => {
    vi.restoreAllMocks();
  });

  it('submits a valid calculation and displays the result', async () => {
    const fetchMock = vi.spyOn(globalThis, 'fetch').mockResolvedValue(
      new Response(JSON.stringify({ operation: 'add', result: 8 }), {
        status: 200,
        headers: { 'Content-Type': 'application/json' },
      }),
    );

    render(<App />);

    await userEvent.type(screen.getByLabelText(/first number/i), '3');
    await userEvent.type(screen.getByLabelText(/second number/i), '5');
    await userEvent.click(screen.getByRole('button', { name: /calculate/i }));

    expect(await screen.findByText('8')).toBeInTheDocument();
    expect(fetchMock).toHaveBeenCalledWith(
      'http://localhost:8080/api/calculate',
      expect.objectContaining({
        method: 'POST',
        body: JSON.stringify({ operation: 'add', a: 3, b: 5 }),
      }),
    );
  });

  it('validates missing second operand before calling the API', async () => {
    const fetchMock = vi.spyOn(globalThis, 'fetch');

    render(<App />);

    await userEvent.type(screen.getByLabelText(/first number/i), '3');
    await userEvent.click(screen.getByRole('button', { name: /calculate/i }));

    expect(await screen.findByText(/valid second number/i)).toBeInTheDocument();
    expect(fetchMock).not.toHaveBeenCalled();
  });

  it('uses one operand for square root', async () => {
    const fetchMock = vi.spyOn(globalThis, 'fetch').mockResolvedValue(
      new Response(JSON.stringify({ operation: 'sqrt', result: 4 }), {
        status: 200,
        headers: { 'Content-Type': 'application/json' },
      }),
    );

    render(<App />);

    await userEvent.click(screen.getByRole('button', { name: /square root/i }));
    await userEvent.type(screen.getByLabelText(/^number$/i), '16');
    await userEvent.click(screen.getByRole('button', { name: /calculate/i }));

    expect(screen.queryByLabelText(/second number/i)).not.toBeInTheDocument();
    expect(await screen.findByText('4')).toBeInTheDocument();
    expect(fetchMock).toHaveBeenCalledWith(
      'http://localhost:8080/api/calculate',
      expect.objectContaining({
        body: JSON.stringify({ operation: 'sqrt', a: 16 }),
      }),
    );
  });

  it('displays backend errors', async () => {
    vi.spyOn(globalThis, 'fetch').mockResolvedValue(
      new Response(JSON.stringify({ error: 'division by zero' }), {
        status: 400,
        headers: { 'Content-Type': 'application/json' },
      }),
    );

    render(<App />);

    await userEvent.click(screen.getByRole('button', { name: /division/i }));
    await userEvent.type(screen.getByLabelText(/first number/i), '3');
    await userEvent.type(screen.getByLabelText(/second number/i), '0');
    await userEvent.click(screen.getByRole('button', { name: /calculate/i }));

    expect(await screen.findByText(/division by zero/i)).toBeInTheDocument();
  });
});

