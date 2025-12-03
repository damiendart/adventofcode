# This file was written by Damien Dart, <damiendart@pobox.com>. This is
# free and unencumbered software released into the public domain. For
# more information, please refer to the accompanying "UNLICENCE" file.

(def input-parser (peg/compile ~{
  :main (split "," (group :range))
  :range (sequence (number :d+) "-" (number :d+))
}))

(def invalid-id (peg/compile ~{
  :main (choice
    (* (<- (repeat 5 :d)) (some (backmatch)) -1)
    (* (<- (repeat 4 :d)) (some (backmatch)) -1)
    (* (<- (repeat 3 :d)) (some (backmatch)) -1)
    (* (<- (repeat 2 :d)) (some (backmatch)) -1)
    (* (<- (repeat 1 :d)) (some (backmatch)) -1)
  )
}))

(defn solve [input]
  (var sum 0)
  (map
    (fn [range]
      (for i (get range 0) (+ (get range 1) 1)
        (if (peg/match invalid-id (string/format "%d" i)) (set sum (+ sum i)))
      )
    )
    (peg/match input-parser input)
  )
  sum
)

(defn main [&] (->> (file/read stdin :all) string/trim solve pp))
