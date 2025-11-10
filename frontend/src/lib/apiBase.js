// src/lib/apiBase.js
function trimSlashes(s = '') {
    return String(s).replace(/\/+$/,'').replace(/^\/+/,'');
}

const runtime = typeof window !== 'undefined' ? window.__ENV?.API_BASE : undefined;
const fromVite = import.meta?.env?.VITE_API_BASE;

// Regla: API_BASE debe TERMINAR en /api (lo inyectamos así en Azure).
export const API_BASE = (runtime || fromVite || '/api').replace(/\/+$/,''); // sin slash final

// Si pasás path, concatena; si no, devuelve la base
export function apiUrl(path = '') {
    if (!path) return API_BASE;
    return `${API_BASE}/${trimSlashes(path)}`;
}
