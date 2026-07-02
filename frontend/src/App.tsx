import { FormEvent, useMemo, useState } from 'react';
import { calculate, type Operation } from './api';
import { getOperation, operations } from './operations';

type FormState = {
  operation: Operation;
  a: string;
  b: string;
};

const initialState: FormState = {
  operation: 'add',
  a: '',
  b: '',
};

export function App() {
  const [form, setForm] = useState<FormState>(initialState);
  const [result, setResult] = useState<number | null>(null);
  const [error, setError] = useState('');
  const [isSubmitting, setIsSubmitting] = useState(false);

  const selectedOperation = useMemo(() => getOperation(form.operation), [form.operation]);

  async function handleSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();
    setError('');
    setResult(null);

    const firstOperand = Number(form.a);
    const secondOperand = Number(form.b);

    if (form.a.trim() === '' || Number.isNaN(firstOperand)) {
      setError('Enter a valid first number.');
      return;
    }

    if (selectedOperation.needsSecondOperand && (form.b.trim() === '' || Number.isNaN(secondOperand))) {
      setError('Enter a valid second number.');
      return;
    }

    setIsSubmitting(true);
    try {
      const payload = selectedOperation.needsSecondOperand
        ? { operation: form.operation, a: firstOperand, b: secondOperand }
        : { operation: form.operation, a: firstOperand };
      const response = await calculate(payload);
      setResult(response.result);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Calculation failed.');
    } finally {
      setIsSubmitting(false);
    }
  }

  function updateField<Key extends keyof FormState>(key: Key, value: FormState[Key]) {
    setForm((current) => ({ ...current, [key]: value }));
  }

  return (
    <main className="app-shell">
      <section className="calculator" aria-labelledby="calculator-title">
        <div className="calculator-header">
          <p className="eyebrow">Sezzle assignment</p>
          <h1 id="calculator-title">Calculator</h1>
        </div>

        <form className="calculator-form" onSubmit={handleSubmit}>
          <fieldset className="operation-grid" aria-label="Operation">
            {operations.map((operation) => (
              <button
                aria-label={operation.label}
                aria-pressed={operation.value === form.operation}
                className="operation-button"
                key={operation.value}
                onClick={() => updateField('operation', operation.value)}
                title={operation.label}
                type="button"
              >
                {operation.symbol}
              </button>
            ))}
          </fieldset>

          <label className="field">
            <span>{selectedOperation.needsSecondOperand ? 'First number' : 'Number'}</span>
            <input
              inputMode="decimal"
              onChange={(event) => updateField('a', event.target.value)}
              placeholder="0"
              type="number"
              value={form.a}
            />
          </label>

          {selectedOperation.needsSecondOperand ? (
            <label className="field">
              <span>{form.operation === 'percentage' ? 'Percent' : 'Second number'}</span>
              <input
                inputMode="decimal"
                onChange={(event) => updateField('b', event.target.value)}
                placeholder="0"
                type="number"
                value={form.b}
              />
            </label>
          ) : null}

          <button className="submit-button" disabled={isSubmitting} type="submit">
            {isSubmitting ? 'Calculating' : 'Calculate'}
          </button>
        </form>

        <output aria-live="polite" className="result-panel">
          {error ? <span className="error">{error}</span> : null}
          {result !== null ? (
            <>
              <span className="result-label">Result</span>
              <strong>{formatResult(result)}</strong>
            </>
          ) : null}
          {!error && result === null ? <span className="muted">Ready</span> : null}
        </output>
      </section>
    </main>
  );
}

function formatResult(value: number): string {
  return Number.isInteger(value) ? value.toString() : value.toFixed(8).replace(/0+$/, '').replace(/\.$/, '');
}

