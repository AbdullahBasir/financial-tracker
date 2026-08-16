// src/lib/format.ts
export function formatDate(isoString: string): string {
  return new Date(isoString).toLocaleDateString('en-GB', {
    year: 'numeric',
    month: 'short',
    day: 'numeric',
  })
}
