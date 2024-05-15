import axios from 'axios';

export const connect = () => {
	document.cookie = 'X-Authorization=user-id-1';
	const ws = new WebSocket('ws://localhost:8080/ws');

	ws.addEventListener('message', (message: unknown) => {
		console.log(message);
	});
};

export const createGame = async () => {
	const host = 'http://localhost:8080';
	const endpoint = '/api/v1/game';

	const res = await axios.post(host + endpoint);
	console.log(res.data);
};
