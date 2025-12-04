# This file was written by Damien Dart, <damiendart@pobox.com>. This is
# free and unencumbered software released into the public domain. For
# more information, please refer to the accompanying "UNLICENCE" file.

(def input-parser (peg/compile ~(* (some (* (<- (set "@.")) (any :s))))))

(defn solve [input]
  (var diagram (peg/match input-parser input))
  (var removed-rolls 0)

  (def width (string/find "\n" input))
  (def height (/ (length diagram) width))

  (var h nil)
  (while (not= h 0)
    (var i 0)
    (set h 0)

    (while (< i (length diagram))
      (var adjacent-roll-count 0)
      (var column-idx (mod i width))
      (var row-idx (math/floor (/ i height)))

      (if (and (not= row-idx 0) (not= column-idx 0))
        (if (= (get diagram (- i width 1)) "@") (++ adjacent-roll-count))
      )

      (if (not= row-idx 0)
        (if (= (get diagram (- i width)) "@") (++ adjacent-roll-count))
      )

      (if (and (not= row-idx 0) (not= column-idx (- width 1)))
        (if (= (get diagram (- i width -1)) "@") (++ adjacent-roll-count))
      )

      (if (not= column-idx 0)
        (if (= (get diagram (- i 1)) "@") (++ adjacent-roll-count))
      )

      (if (not= column-idx (- width 1))
        (if (= (get diagram (+ i 1)) "@") (++ adjacent-roll-count))
      )

      (if (and (not= row-idx (- height 1)) (not= column-idx 0))
        (if (= (get diagram (+ i width -1)) "@") (++ adjacent-roll-count))
      )

      (if (not= row-idx (- height 1))
        (if (= (get diagram (+ i width)) "@") (++ adjacent-roll-count))
      )

      (if (and (not= row-idx (- height 1)) (not= column-idx (- width 1)))
        (if (= (get diagram (+ i width 1)) "@") (++ adjacent-roll-count))
      )

      (if (and (= (get diagram i) "@") (< adjacent-roll-count 4)) (do
        (++ h)
        (++ removed-rolls)
        (set (diagram i) "X")))

      (++ i)
    )
  )

  removed-rolls
)

(defn main [&] (->> (file/read stdin :all) string/trim solve pp))
