import type { Operation } from './api';

export type OperationConfig = {
  value: Operation;
  label: string;
  symbol: string;
  needsSecondOperand: boolean;
};

export const operations: OperationConfig[] = [
  { value: 'add', label: 'Addition', symbol: '+', needsSecondOperand: true },
  { value: 'subtract', label: 'Subtraction', symbol: '-', needsSecondOperand: true },
  { value: 'multiply', label: 'Multiplication', symbol: 'x', needsSecondOperand: true },
  { value: 'divide', label: 'Division', symbol: '/', needsSecondOperand: true },
  { value: 'power', label: 'Exponent', symbol: '^', needsSecondOperand: true },
  { value: 'sqrt', label: 'Square root', symbol: 'sqrt', needsSecondOperand: false },
  { value: 'percentage', label: 'Percentage', symbol: '%', needsSecondOperand: true },
];

export function getOperation(value: Operation): OperationConfig {
  return operations.find((operation) => operation.value === value) ?? operations[0];
}

