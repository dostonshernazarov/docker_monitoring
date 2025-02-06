import React, { useEffect, useState } from 'react';

const App: React.FC = () => {
	const [containers, setContainers] = useState<{ ip: string; lastPing: string }[]>([]);

	useEffect(() => {
		fetch("http://backend:8080/containers")
			.then(res => res.json())
			.then(data => setContainers(data))
			.catch(err => console.error(err));
	}, []);

	return (
		<div className="container">
			<h1>Docker Monitoring</h1>
			<table>
				<thead>
					<tr>
						<th>IP Address</th>
						<th>Last Ping</th>
					</tr>
				</thead>
				<tbody>
					{containers.map((c, i) => (
						<tr key={i}>
							<td>{c.ip}</td>
							<td>{c.lastPing}</td>
						</tr>
					))}
				</tbody>
			</table>
		</div>
	);
};

export default App;