INSERT INTO `items` (`item_name`, `item_disp_name`) VALUES 
('salmon', 'Salmon'),
('tuna', 'Tuna'),
('prawn', 'Prawn'),
('lobster', 'Lobster'),
('crab', 'Crab'),
('squid', 'Squid'),
('scallop', 'Scallop'),
('mackerel', 'Mackerel');

INSERT INTO `prices` (`item_id`, `rate`, `is_default`) VALUES 
((SELECT id FROM items WHERE item_name = 'salmon'), 300.00, 1),
((SELECT id FROM items WHERE item_name = 'tuna'), 250.00, 1),
((SELECT id FROM items WHERE item_name = 'prawn'), 200.00, 1),
((SELECT id FROM items WHERE item_name = 'lobster'), 500.00, 1),
((SELECT id FROM items WHERE item_name = 'crab'), 400.00, 1),
((SELECT id FROM items WHERE item_name = 'squid'), 150.00, 1),
((SELECT id FROM items WHERE item_name = 'scallop'), 350.00, 1),
((SELECT id FROM items WHERE item_name = 'mackerel'), 100.00, 1);
