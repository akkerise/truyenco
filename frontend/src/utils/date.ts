import dayjs from 'dayjs';

export const formatDate = (d: Date) => dayjs(d).format('YYYY-MM-DD');
