-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
INSERT INTO compliments (text) VALUES ('%s, ты самый лучший человек на планете');
INSERT INTO compliments (text) VALUES ('%s, ты прекрасный человечек');
INSERT INTO compliments (text) VALUES ('%s, у тебя все получится!');
INSERT INTO compliments (text) VALUES ('%s, улыбнись!');
INSERT INTO compliments (text) VALUES ('%s, твоя улыбка светит, как солнце"');
INSERT INTO compliments (text) VALUES ('%s, этому миру нужны такие люди, как ты"');
INSERT INTO compliments (text) VALUES ('%s, ты неповторимый и совершенный человечек');
INSERT INTO compliments (text) VALUES ('%s, ты вызываешь восхищение');
INSERT INTO compliments (text) VALUES ('%s, твоя добрая энергия притягивает всех вокруг');
INSERT INTO compliments (text) VALUES ('%s, ты источник вдохновения');
INSERT INTO compliments (text) VALUES ('%s, никогда не сдавайся');
INSERT INTO compliments (text) VALUES ('%s, невозможно быть милее, чем ты!');
INSERT INTO compliments (text) VALUES ('%s, потрясающе умеешь слушать и поддерживать!');
INSERT INTO compliments (text) VALUES ('%s, c тобой каждый день становится особенным');
INSERT INTO compliments (text) VALUES ('%s, ты очень необычный и уникальный человек');
INSERT INTO compliments (text) VALUES ('%s, прекрасного тебе дня!');
INSERT INTO compliments (text) VALUES ('%s, замечательного тебе дня!');
INSERT INTO compliments (text) VALUES ('%s, ты всех уделаешь сегодня!');
INSERT INTO compliments (text) VALUES ('%s, у тебя потрясающие причёска и улыбка!');
INSERT INTO compliments (text) VALUES ('%s, ты удивительный и прекрасный человек.');
INSERT INTO compliments (text) VALUES ('%s, ты  воплощение  мечты');
INSERT INTO compliments (text) VALUES ('%s, все будет прекрасно, улыбайся!');
INSERT INTO compliments (text) VALUES ('%s, ты мотивируешь  быть лучшим');
INSERT INTO compliments (text) VALUES ('%s, ты просто в фантастической форме!');
INSERT INTO compliments (text) VALUES ('%s, никто на свете не сравнится с тобой.');
INSERT INTO compliments (text) VALUES ('%s, ты превосходный человечек');
INSERT INTO compliments (text) VALUES ('%s, в тебя невозможно не влюбиться');
INSERT INTO compliments (text) VALUES ('%s, ты просто блистательный человечек!');
INSERT INTO compliments (text) VALUES ('%s, твои чарующие глаза притягивают как магниты');
INSERT INTO compliments (text) VALUES ('%s, ну давай же, надери им всем зад!');
INSERT INTO compliments (text) VALUES ('%s, ты — невероятно талантливая и креативная личность');
INSERT INTO compliments (text) VALUES ('%s, твой смех — самая сладкая мелодия');
INSERT INTO compliments (text) VALUES ('%s, ты  умеешь создавать магию в повседневных моментах');
INSERT INTO compliments (text) VALUES ('%s, ты словно магия, оживляющая серые будни');
INSERT INTO compliments (text) VALUES ('%s, ты наполняешь мир красотой и смыслом');
INSERT INTO compliments (text) VALUES ('%s, тебя ждет сегодня успех');
INSERT INTO compliments (text) VALUES ('%s, все будет хорошо!');
INSERT INTO compliments (text) VALUES ('%s, интересного, насыщенного, яркого дня! Приветливых улыбок и хорошего настроения! ');
INSERT INTO compliments (text) VALUES ('%%, улыбнись, ты оч красивый человек"');
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
TRUNCATE compliments;
-- +goose StatementEnd
