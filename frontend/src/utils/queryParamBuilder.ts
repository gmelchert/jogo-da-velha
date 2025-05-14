export const queryParamBuilder = (query?: Record<string, any>) => {
    if (!query) return '';

    const queryParams = new URLSearchParams();

    Object.entries(query).forEach(([key, value]) => {
        if (value !== undefined && value !== null) {
            queryParams.append(key, String(value));
        }
    });

    return queryParams.toString();
}